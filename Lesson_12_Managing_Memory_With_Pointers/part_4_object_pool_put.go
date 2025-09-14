func (p *ObjectPool) Put(obj *Object) {
	obj.ID = 0 // Resetting object state
	select {
	case p.pool <- obj:
	default:
		// Pool is full, object is discarded
	}
}
