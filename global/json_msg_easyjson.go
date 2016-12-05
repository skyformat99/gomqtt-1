// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package global

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal(in *jlexer.Lexer, out *JsonMsg) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "facc":
			if in.IsNull() {
				in.Skip()
				out.FAcc = nil
			} else {
				out.FAcc = in.Bytes()
			}
		case "ftopic":
			if in.IsNull() {
				in.Skip()
				out.FTopic = nil
			} else {
				out.FTopic = in.Bytes()
			}
		case "type":
			out.Type = int(in.Int())
		case "qos":
			out.Qos = int(in.Int())
		case "time":
			out.Time = int(in.Int())
		case "nick":
			if in.IsNull() {
				in.Skip()
				out.Nick = nil
			} else {
				out.Nick = in.Bytes()
			}
		case "mi":
			if in.IsNull() {
				in.Skip()
				out.MsgID = nil
			} else {
				out.MsgID = in.Bytes()
			}
		case "m":
			if in.IsNull() {
				in.Skip()
				out.Msg = nil
			} else {
				out.Msg = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal(out *jwriter.Writer, in JsonMsg) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"facc\":")
	out.Base64Bytes(in.FAcc)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ftopic\":")
	out.Base64Bytes(in.FTopic)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	out.Int(int(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"qos\":")
	out.Int(int(in.Qos))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"time\":")
	out.Int(int(in.Time))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nick\":")
	out.Base64Bytes(in.Nick)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mi\":")
	out.Base64Bytes(in.MsgID)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"m\":")
	out.Base64Bytes(in.Msg)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JsonMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JsonMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JsonMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JsonMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal(l, v)
}
func easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal1(in *jlexer.Lexer, out *Messages) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "compress":
			out.Compress = int(in.Int())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				out.Data = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal1(out *jwriter.Writer, in Messages) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"compress\":")
	out.Int(int(in.Compress))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"data\":")
	out.Base64Bytes(in.Data)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Messages) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Messages) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA390b27cEncodeGithubComAiyunGomqttGlobal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Messages) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Messages) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA390b27cDecodeGithubComAiyunGomqttGlobal1(l, v)
}