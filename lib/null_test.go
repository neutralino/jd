package jd

import (
	"testing"
)

func TestNullJson(t *testing.T) {
	ctx := newTestContext(t)
	checkJson(ctx, `null`, `null`)
}

func TestNullEqual(t *testing.T) {
	ctx := newTestContext(t)
	checkEqual(ctx, `null`, `null`)
}

func TestNullNotEqual(t *testing.T) {
	ctx := newTestContext(t)
	checkNotEqual(ctx, `null`, `0`)
	checkNotEqual(ctx, `null`, `[]`)
	checkNotEqual(ctx, `null`, `{}`)
}

func TestNullHash(t *testing.T) {
	ctx := newTestContext(t)
	checkHash(ctx, `null`, `null`, true)
	checkHash(ctx, `null`, ``, false)
}

func TestNullDiff(t *testing.T) {
	ctx := newTestContext(t)
	checkDiff(ctx, `null`, `null`)
	checkDiff(ctx, `null`, ``,
		`@ []`,
		`- null`)
	checkDiff(ctx, ``, `null`,
		`@ []`,
		`+ null`)
	ctx = ctx.withMetadata(MERGE)
	checkDiff(ctx, `true`, `null`,
		`@ [["MERGE"]]`,
		`+ null`)
	checkDiff(ctx, `null`, `true`,
		`@ [["MERGE"]]`,
		`+ true`)
}

func TestNullPatch(t *testing.T) {
	ctx := newTestContext(t)
	checkPatch(ctx, `null`, `null`)
	checkPatch(ctx, `null`, ``,
		`@ []`,
		`- null`)
	checkPatch(ctx, ``, `null`,
		`@ []`,
		`+ null`)
	// Null deletes a node
	checkPatch(ctx, `null`, ``,
		`@ [["MERGE"]]`,
		`+ null`)
}

func TestNullPatchError(t *testing.T) {
	ctx := newTestContext(t)
	checkPatchError(ctx, `null`,
		`@ []`,
		`- 0`)
}
