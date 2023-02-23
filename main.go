package main

type Data struct {
	Value uint8 // (uint8 == 1 byte)
	// isTrue       bool //0 => math.Pow(2, 0) => 1
	// isClass      bool //1 => math.Pow(2, 1) => 2
	// isLive       bool //2 => math.Pow(2, 2) => 4
	// isIron       bool //3 => math.Pow(2, 3) => 8
	// isSearchable bool //4 => math.Pow(2, 4) => 16
	// isCode       bool //5 => math.Pow(2, 5) => 32
	// isFlag       bool //6 => math.Pow(2, 6) => 64
	// isFunc       bool //7 => math.Pow(2, 7) => 128
}

func (d *Data) getIsIron() bool {
	v := d.Value >> 3
	return v&1 == 1
}

func (d *Data) setIsIron() {
	d.Value |= 8
}

func (d *Data) getIsCode() bool {
	v := d.Value >> 5
	return v&1 == 1
}

func (d *Data) setIsCode() {
	d.Value |= 32
}

func (d *Data) getIsFlag() bool {
	v := d.Value >> 6
	return v&1 == 1
}

func main() {
	d := Data{Value: 0}   // d.Value is 0000 0000 => 0
	d.setIsIron()         // d.Value is 0000 1000 => 8
	d.setIsCode()         // d.Value is 0010 1000 => 40
	iron := d.getIsIron() // 0010 1000 >> 5 => (0000 0001 & 0000 0001) = 1 => true
	println(iron)
	code := d.getIsCode() // 0010 1000 >> 3 => (0000 0101 & 0000 0001) = 1 => true
	println(code)
	flag := d.getIsFlag() // 0010 1000 >> 6 => (0000 0000 & 0000 0001) = 0 => false
	println(flag)
}

//0001 0001    >> 0   //Default 			=> 17
//0000 1000    >> 1   //If shift 1 time 	=> 8
//0000 0100    >> 2   //If shift 2 times 	=> 4
//0000 0010    >> 3   //If shift 3 times    => 2
//0000 0001    >> 4   //If shift 4 times    => 1
