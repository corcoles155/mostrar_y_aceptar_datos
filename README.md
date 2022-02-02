# Mostrar y aceptar datos en GO

En GOLANG podemos usar la funcion "Scan" del paquete "fmt" para aceptar datos ```fmt.Scan("%d", &numero1)```. Esta función tiene un bug en Windows, si por ejemplo ponemos dos entradas de datos ```fmt.Scan("%d", &numero1)``` y ```fmt.Scan("%d", &numero2)```, escribimos el valor "1" para la variable "numero1" y damos a enter va a tomar que el valor de la variable "numero2" es "\r". Esto se debe a que en Linux el salto de línea es un unico byte ("\n") y para Windows son dos  ("\n\r"). Esto se resuelve utilizando "Scanln" en lugar de "Scan" ```fmt.Scanln("%d", &numero1)```.

Utilizamos el paquete "bufio" https://pkg.go.dev/bufio, el paquete bufio implementa E/S con buffer.Envuelve un objeto io.Reader o io.Writer,creando otro objeto (Reader o Writer)que también implementa la interfaz pero que proporciona un buffer y alguna ayuda para la E/S textual.

Utilizamos el paquete "os" https://pkg.go.dev/os@go1.17.6 para leer de la entrada estándar (teclado). ```os.Stdin```
