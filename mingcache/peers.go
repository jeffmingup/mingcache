package mingcache

import mingchachepd "mingcache/mingcachepd"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}
type PeerGetter interface {
	Get(in *mingchachepd.Request, out *mingchachepd.Response) error
}
