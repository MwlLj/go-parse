package url

import (
    "bytes"
)

type Param struct {
    Key bytes.Buffer
    Value bytes.Buffer
}

func (self *Param) clear() {
    self.Key = bytes.Buffer{}
    self.Value = bytes.Buffer{}
}

type Url struct {
    Proto bytes.Buffer
    Addr bytes.Buffer
    Path bytes.Buffer
    Params []Param
}

type urlSplitMode int8
const (
    _ urlSplitMode = iota
    urlSplitModeProto
    urlSplitModeAddr
    urlSplitModePath
    urlSplitModeParam
)

type urlSplitKvMode int8
const (
    _ urlSplitKvMode = iota
    urlSplitKvModeNormal
    urlSplitKvModeKey
    urlSplitKvModeValue
)

func UrlSplit(url string) *Url {
    u := Url{}
    m := urlSplitModeProto
    kvm := urlSplitKvModeNormal
    ib := 0
    param := Param{}
    param.clear()
    for _, c := range url {
        switch m {
        case urlSplitModeProto:
            if ib == 2 && c == '/' {
                ib = 0
                m = urlSplitModeAddr
                continue
            }
            if c != ':' && c != '/' {
                u.Proto.WriteRune(c)
            } else {
                ib += 1
            }
        case urlSplitModeAddr:
            if c == '/' {
                u.Path.WriteRune('/')
                m = urlSplitModePath
                continue
            } else {
                u.Addr.WriteRune(c)
            }
        case urlSplitModePath:
            if c == '?' {
                m =  urlSplitModeParam
                kvm = urlSplitKvModeKey
                continue
            } else {
                u.Path.WriteRune(c);
            }
        case urlSplitModeParam:
            switch kvm {
            // case urlSplitKvModeNormal:
            case urlSplitKvModeKey:
                if c == '=' {
                    kvm = urlSplitKvModeValue
                } else {
                    param.Key.WriteRune(c)
                }
            case urlSplitKvModeValue:
                if c == '&' {
                    u.Params = append(u.Params, param)
                    param.clear()
                    kvm = urlSplitKvModeKey
                } else {
                    param.Value.WriteRune(c);
                }
            }
        }
    }
    switch kvm {
    case urlSplitKvModeNormal:
    default:
        u.Params = append(u.Params, param)
    }
    return &u
}
