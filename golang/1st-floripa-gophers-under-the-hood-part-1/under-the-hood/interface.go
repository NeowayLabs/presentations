type Stringer interface {
    String() string;
}

type Binary uint64

func (i Binary) Get() uint64 {
    return uint64(i)
}

func (i Binary) String() string {
    return strconv.Uitob64(i.Get(), 2)
}

