package iaas

//go:generate proteus -p github.com/Serabe/iaas/iaas -f $GOPATH/src/github.com/Serabe/iaas/iaas/protos

//proteus:generate
func Inc(i int64) int64 {
	return i + 1
}
