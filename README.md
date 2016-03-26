A.LAPORAN

1.Scope dalam golang adalah cara penulisan variabel yang diletakkan didalam suatu fungsi tertentu, dimana variabel tersebut tidak dapat dipanggil oleh fungsi yang lain. Kalau dalam java masuk dalam pembahasan hak akses yaitu perti private, public, dan default. Dalam menentukan scope dalam golang yaitu dengan menaruh variabel yang ingin diakses oleh banyak fungsi (seperti public) yaitu dengan meletakkan variabel dipaling atas/diluar dari fungsi:

--> Variabel adalah nama yang mewakili suatu nilai, isi dari variabel dapat diubah --> Costanta adalah variabel diman nilanya tidak dapat diubah/tetap
B.TUGAS

1.Script program (Membuat Program sederhana dengan golang tentang mengkonfersi niali dengan golang):

 package main
 import "fmt"
     func main(){
       var f float32
       fmt.Print("Masukan Nilai Fahrenheit = ")
       fmt.Scanf("%f", &f)
       var c float32
       c= (f-32)*5/9
       fmt.Println("Hasil Konfersi ke-Celcius = ", c)
   }
2.Script program (Mengkonversi Matauang Dolar ke Rupaiah dengan golang):

package main
import "fmt"
func main(){
    var d float32
    fmt.Print("Masukan Mata Uang Dolar = ")
    fmt.Scanf("%f", &d)

    var r float32

    r= d*13510
    fmt.Println("Hasil Konfersi Mata Uang Rupiah = Rp.", r)
