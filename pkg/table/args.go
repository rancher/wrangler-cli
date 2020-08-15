package table

type Args struct {
	Quiet  bool   `name:"quiet" usage:"only print ID" short:"q"`
	Format string `name:"format" usage:"format(yaml/json/jsoncompact/raw)"`
}
