# First steps in Golang - Workshop - Mercado Libre

<a href="https://www.google.com.ar/imgres?imgurl=https%3A%2F%2Fcamo.githubusercontent.com%2Ff04a7348c3f516c528405b48eccc96d3ad1abdeb%2F68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722d736964655f636f6c6f722e706e67&imgrefurl=https%3A%2F%2Fgithub.com%2Fgolang-samples%2Fgopher-vector&docid=VsdewyBT_3vWLM&tbnid=Mf6d7krXOijd7M%3A&vet=10ahUKEwiYxfWZk5zXAhXRl5AKHdBdBPsQMwhDKBUwFQ..i&w=416&h=554&bih=759&biw=1440&q=golang%20logo&ved=0ahUKEwiYxfWZk5zXAhXRl5AKHdBdBPsQMwhDKBUwFQ&iact=mrc&uact=8"/>

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
