package vms


/*
	约定参数：
	[]
 */
func CreateVM(args []string) error {

	if err := execCreate(); err != nil {
		return err
	}

	return nil
}

func execCreate() error {

	return nil
}
