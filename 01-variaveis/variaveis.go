package main

import "fmt"

var linguagem string
var soma bool

func main() {
	// declaracao explicita do tipo
	var msg string
	var n1, n2 int
	var n3, n4 = 15, 25

	//declaracao curta
	msg2 := "ola,mudo"

	msg = "oi"
	n1 = 10
	n2 = 20

	fmt.Println(msg, n1, n2, n3, n4, msg2)

	f2()

	var numero float64
	numero = 4.7

	var n6 complex128

	const Pi = 3,14

}

/*
inteiros

int8	128 a 127
uint8	
int16	32768 a 32.767
uint16
int32
uint32
int64
uint64

byte	uint8
rune int32 caracters unicode

int aqui = int32 ou int64, dependendo da arquittura do computador


------------------------------

flutuantes

float32	precisao de 6 a 9 digitos
float64 padrao 	 precisao de 15 a 17 digitos

é necessario epecificar o tamanho do float
caso a declaracao seja implicita, é assunto float64

-----------------------------------

complexo

complex64
complex128

-----------------------------------

string
bool 	true or false - minusculo tudo sempre

nill	ausencia de dados 



*/
