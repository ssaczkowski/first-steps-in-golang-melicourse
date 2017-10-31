# first-steps-in-golang-melicourse


- Ejercicio 1: 
Hacer el mitico -> "Hello world" en Golang.

- Ejercicio 2: 
Hacer un programa que reciba un parámetro p por consola y que cacule la suma de los valores v tales que 0 <= v <= p , siendo v un número divisible por 3 o por 5
Ayuda para leer argumentos desde la consola
for i := 1; i < len(os.Args); i++ {
    fmt.Println(os.Args[i])
}

- Ejercicio 3: 
(struct, punteros, slices) Hacer un programa que reciba por parametro una lista de enteros, los almacene en un slice, popule un arbol binario y despues imprima los valores en orden ascendente.

- Ejercicio 4: 
Modelar la funcionalidad para un sistema de cines que calcule los ingresos netos de una función en base a los asistentes y al precio base de la entrada. Existen tres tipos de asistentes y tienen las siguientes características:
Normales: Pagan el 100%
Jubilados: Tienen un 50% de descuento
Invitados: No pagan nada

- Ejercicio 5: 
Agregar un test, un benchmark y un example al ejercicio de suma de un valor parametrizado p 
entre 0 <= v <= p.  1º Ejercicio del curso.

- Ejercicio 6: Evolucionar el primer ejercicio del curso de forma tal que ahora se ejecute 
en un webserver que escuche en el puerto 8080 y sea invocado al llamar al servicio 
GET /calculate?value=p. El servicio debe devolver la respuesta en formato JSON
