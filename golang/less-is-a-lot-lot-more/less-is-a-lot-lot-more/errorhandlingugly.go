_, err = fd.Write(p0[a:b])
if err != nil {
	return err
}
_, err = fd.Write(p1[c:d])
if err != nil {
	return err
}
_, err = fd.Write(p2[e:f])
if err != nil {
	return err
}
// and so on
