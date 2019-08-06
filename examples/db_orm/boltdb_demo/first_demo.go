package main

import (
    "fmt"
    "github.com/boltdb/bolt"
    "log"
    "time"
)

func main() {
    db, err := bolt.Open("mybolt.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // read-write transaction
    //err = db.Update(func(tx *bolt.Tx) error {
    //    // ...
    //    return nil
    //})
    // batch read-write transaction
    //err := db.Batch(func(tx *bolt.Tx) error {
    //
    //    return nil
    //})

    // read-only transaction
    //err = db.View(func(tx *bolt.Tx) error {
    //    // ...
    //    return nil
    //})

    if err := write(db); err != nil {
        log.Fatal(err)
    }

    _, err = read(db)
    if err != nil {
        log.Fatal(err)
    }

    return
}

func write(db *bolt.DB) error {
    // manually transaction
    tx, err := db.Begin(true)
    if err != nil {
        log.Fatal(err)
    }
    // 后面有commit，commit成功后再rollback就没效果了
    defer tx.Rollback()

    bk, err := tx.CreateBucket([]byte("MyBucket"))
    if err != nil {
        return err
    }

    if err := bk.Put([]byte("say"), []byte("bonjour monte")); err != nil {
        //log.Fatal(err)
        return err
    }

    // commit
    if err := tx.Commit(); err != nil {
        //log.Fatal(err)
        return err
    }

    return nil
}

func read(db *bolt.DB) (string, error) {
    var ret string
    err := db.View(func(tx *bolt.Tx) error {
        bk := tx.Bucket([]byte("MyBucket"))
        bt := bk.Get([]byte("say"))
        fmt.Println(">>>", string(bt))
        ret = string(bt)
        return nil
    })
    if err != nil {
        return "", err
    }

    return ret, nil
}
