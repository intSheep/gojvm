package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gojvm/common/classfile"
	"gojvm/common/instructions"
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	// 从codeAttribute获取运行所需最大栈深度、局部变量所需执行空间
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		logrus.Errorf("localVars:%v\n", frame.LocalVars())
		logrus.Errorf("OperandStack:%v\n", frame.LocalVars())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := base.NewBytecodeReader()
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// 解码指令
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// 执行指令
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
