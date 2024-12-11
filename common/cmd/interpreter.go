package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gojvm/common/instructions"
	"gojvm/common/instructions/base"
	"gojvm/common/rtda"
	"gojvm/common/rtda/heap"
)

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
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
