package jd

type jsonBool bool

var _ JsonNode = jsonBool(true)

func (b jsonBool) Json(metadata ...Metadata) string {
	return renderJson(b.raw(metadata))
}

func (b jsonBool) Yaml(metadata ...Metadata) string {
	return renderYaml(b.raw(metadata))
}

func (b jsonBool) raw(metadata []Metadata) interface{} {
	return bool(b)
}

func (b1 jsonBool) Equals(n JsonNode, metadata ...Metadata) bool {
	b2, ok := n.(jsonBool)
	if !ok {
		return false
	}
	return b1 == b2
}

func (b jsonBool) hashCode(metadata []Metadata) [8]byte {
	if b {
		return [8]byte{0x24, 0x6B, 0xE3, 0xE4, 0xAF, 0x59, 0xDC, 0x1C} // Random bytes
	} else {
		return [8]byte{0xC6, 0x38, 0x77, 0xD1, 0x0A, 0x7E, 0x1F, 0xBF} // Random bytes
	}
}

func (b jsonBool) Diff(n JsonNode, metadata ...Metadata) Diff {
	strategy := getPatchStrategy(metadata)
	return b.diff(n, make(path, 0), metadata, strategy)
}

func (b jsonBool) diff(n JsonNode, path path, metadata []Metadata, strategy patchStrategy) Diff {
	d := make(Diff, 0)
	if b.Equals(n) {
		return d
	}
	var e DiffElement
	switch strategy {
	case mergePatchStrategy:
		e = DiffElement{
			Path:      path.prependMetadataMerge(),
			NewValues: nodeList(n),
		}
	default:
		e = DiffElement{
			Path:      path.clone(),
			OldValues: nodeList(b),
			NewValues: nodeList(n),
		}
	}
	return append(d, e)
}

func (b jsonBool) Patch(d Diff) (JsonNode, error) {
	return patchAll(b, d)
}

func (b jsonBool) patch(pathBehind, pathAhead path, oldValues, newValues []JsonNode, strategy patchStrategy) (JsonNode, error) {
	if !pathAhead.isLeaf() {
		return patchErrExpectColl(b, pathAhead[0])
	}
	if len(oldValues) > 1 || len(newValues) > 1 {
		return patchErrNonSetDiff(oldValues, newValues, pathBehind)
	}
	oldValue := singleValue(oldValues)
	newValue := singleValue(newValues)
	switch strategy {
	case mergePatchStrategy:
		if !isVoid(oldValue) {
			return patchErrMergeWithOldValue(pathBehind, oldValue)
		}
		if isNull(newValue) {
			return voidNode{}, nil
		}
	case strictPatchStrategy:
		if !b.Equals(oldValue) {
			return patchErrExpectValue(oldValue, b, pathBehind)
		}
	default:
		return patchErrUnsupportedPatchStrategy(pathBehind, strategy)
	}
	return newValue, nil
}
