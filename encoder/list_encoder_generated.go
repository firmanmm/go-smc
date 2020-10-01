package encoder

//Start of auto generated function for List Encoder

func (l *ListEncoder) EncodeInt(data []int, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeInt8(data []int8, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeInt16(data []int16, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeInt32(data []int32, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeInt64(data []int64, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeUint(data []uint, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeUint8(data []uint8, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeUint16(data []uint16, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeUint32(data []uint32, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeUint64(data []uint64, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeBool(data []bool, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeString(data []string, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeFloat32(data []float32, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeFloat64(data []float64, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) EncodeInterface(data []interface{}, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for _, val := range data {
		if err := l.valueEncoder.Encode(val, writer); err != nil {
			return err
		}
	}
	return nil
}

//End of auto generated function for List Encoder
