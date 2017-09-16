package table

type XmlRow struct {
	items []byte
}

type XmlReader struct {
	rows []*XmlReader
}

func (this *XmlReader) Load(file string) bool {
	return true
}

func (this *XmlReader) Close() {

}
