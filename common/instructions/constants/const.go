package constants

import (
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (a *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (d *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (d *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (f *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (f *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (f *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (i *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (l *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (l *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
