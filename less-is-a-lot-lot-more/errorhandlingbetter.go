var err error
write := func(buf []byte) {
	if err != nil {
		return
	}
	_, err = w.Write(buf)
}
write(p0[a:b])
write(p1[c:d])
write(p2[e:f])
// and so on
if err != nil {
	return err
}
