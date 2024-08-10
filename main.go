package main

import (
	"fmt"
	"os"
	"os/exec"
)

func exit(code int) string {
	//	return fmt.Sprintf("%smov rax, 60\n%smov rdi, %d\n%s%s", INDENT, INDENT, code, INDENT, CALL)
	return INDENT + "mov rax, 60\n" +
		INDENT + fmt.Sprintf("mov rdi, %d\n", code) +
		INDENT + CALL
}

const (
	ENTRY  = "global _start\n_start:\n"
	CALL   = "syscall\n"
	INDENT = "\t"
)

func main() {
	exe_name := "program"
	asm := []byte(ENTRY + exit(8))

	if err := os.WriteFile(exe_name+".asm", asm, 0644); err != nil {
		fmt.Println("Gagal membuat file assembly:", err)
		os.Exit(1)
	}

	assemble := exec.Command("nasm", "-f", "elf64", exe_name+".asm")
	if err := assemble.Run(); err != nil {
		fmt.Println("Gagal membuat file objek:", err)
		os.Exit(1)
	}

	linker := exec.Command("ld", "-o", exe_name, exe_name+".o")
	if err := linker.Run(); err != nil {
		fmt.Println("Gagal membuat file executable:", err)
		os.Exit(1)
	}
}
