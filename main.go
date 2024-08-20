package main

import (
	"fmt"
	"os"
	"os/exec"
	"rangga/lexer"
)

func exit(code int) string {
	return INDENT + "mov rax, 60\n" +
		INDENT + fmt.Sprintf("mov rdi, %d\n", code) +
		INDENT + CALL
}

const (
	ENTRY  = "global _start\n_start:\n"
	CALL   = "syscall\n"
	INDENT = "\t"

	program = "cetak (23 + 2 * 8)"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	switch args[0] {
	case "susun":
		compile("program")
	case "jalan":
		interpret(program)
	}
}

func usage() {
	fmt.Println("Pemakaian: rangga [PERINTAH]")
	fmt.Printf("Perintah tersedia:\n" +
		"susun \t: menyusun kode sumber menjadi program yang dapat dijalankan\n" +
		"jalan \t: menjalankan program tanpa menyusun\n")
}

func compile(exe_name string) {
	asm := []byte(ENTRY + exit(0))

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

func interpret(program string) {
	tokens := lexer.New(program).Tokens()
	//	print(len(tokens))
	for i, token := range tokens {
		fmt.Printf("%d: %v \n", i, token)
	}
}
