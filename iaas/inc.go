package iaas

//go:generate proteus -p github.com/Serabe/iaas/iaas -f $GOPATH/src/github.com/Serabe/iaas/iaas/protos

//proteus:generate
func Inc(i float64) float64 {
	return i + 1
}
