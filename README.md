# Goland-Api
Api para proyectos de recidencias Go

En en este repositorio se encontraran dos ramas la de "proceso" se puede mostrar el trabajo que agrego cada uno y en la rama main se muestra solo el producto final.

para poder ver el trabajo de jenkins sera requerido ir a otro repositorio para que no se confundan los trabajos, ambos realizados por el equipo pero para mantener el orden se realizaron dos repositorios:
link para ir al repode jenkins: https://github.com/CarlosHValadezM1/ApiRestGoJenkins.git

Proceso del docker

lo primero que tenemos que hacer es creado un archivo de nombre Dockerfile(debe usarse exactamente este nombre) Visual estudio identificara este como un docker
![image](https://github.com/user-attachments/assets/fc872b11-5e31-4622-b207-f1837d8cad77)

Dentro de nuestro archivo Dcokerfile tendremos que poner las intrcciones que se llevara acabo, en el cado de la api de go sera con estas
![image](https://github.com/user-attachments/assets/2863db7d-a9a3-4205-9028-5db1a04c4675)

una vez tengamos este archivo de Dockerfile listo, abrimos una nueva terminal en visual y en ella pondremos este comando
![image](https://github.com/user-attachments/assets/f500df79-ba17-4912-8cc1-80dac6b21ea3)

el cual nos permitirar construir nuestra imagen, en caso de que queramos ver nuestras imagenes creadas solo sera requerido usar el comdando Docker images
![image](https://github.com/user-attachments/assets/932aba1f-90c5-4825-9b81-a1b18b99cf7a)

ahora solo nos resta usar el comando Docker run  <nombre_de_tu_imagen> para correr nuestra imagen construida 
![image](https://github.com/user-attachments/assets/8c18c187-a0b4-4ee9-925c-e5cde2ee7e0f)

incluso si se desean ver mas caracteristicas se puede  hacer desde el programa Docker Desktop




#Integrantes del grupo de trabajo:

//Carlos Humberto Valadez Molina
//Mauricio de Jesus Cardona Ramirez
//Chistian Ernesto Silva Pedraza
