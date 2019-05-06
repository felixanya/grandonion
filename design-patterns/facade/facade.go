package facade

import "fmt"

// 数据处理器接口
type IDataProcessor interface {
	Process(string)
}

//
type DataHandler struct {
	decoder 	IDataProcessor
	encoder 	IDataProcessor
	aggregator 	IDataProcessor
}

func NewDataHandler() *DataHandler {
	return &DataHandler{
		decoder: &Decoder{},
		encoder: &Encoder{},
		aggregator: &Aggregator{},
	}
}

func (dh *DataHandler) Decode(data string) {
	dh.decoder.Process(data)
}

func (dh *DataHandler) Encode(data string) {
	dh.encoder.Process(data)
}

func (dh *DataHandler) Aggregate(data string) {
	dh.aggregator.Process(data)
}

// decoder
type Decoder struct {

}

func (d *Decoder) Process(data string) {
	fmt.Printf("decoding data: %s\n", data)
}

// encoder
type Encoder struct {

}

func (e *Encoder) Process(data string) {
	fmt.Printf("encoding data: %s\n", data)
}

// aggregator
type Aggregator struct {

}

func (a *Aggregator) Process(data string) {
	fmt.Printf("aggregating data: %s\n", data)
}
