// composite literal construction
var timeZone = map[string] int {
    "UTC" : 0*60*60,
    "EST" : -5*60*60,
    // and so on
}

// accessing map values
offset := timeZone["EST"]

// checking 0 v.s. non−existence
var seconds int
var ok bool

seconds, ok = timeZone[tz] //comma ok idiom

