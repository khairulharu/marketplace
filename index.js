// 1. Apa yang akan terjadi jika kita menjalankan kode berikut?

for (var index= 0; index <4; index++) {
     setTimeout(()=>{
          console.log(index);
     }, 1000)
}
//A Jelaskan mengapa hasil dari kode tersebut menampilkan angka 4 untuk setiap iterasi, dan
// bagaimana konsep closures berperan dalam situasi ini.
// Kode yang di atas akan mencetak angka 4 sebanyak 4 kali. Ini bukan hasil yang mungkin Anda harapkan, karena Anda mungkin mengharapkan kode untuk mencetak angka 0 sampai 3. Hal ini terjadi karena konsep closure dan event loop dalam JavaScript.

// Mengapa Hasilnya Menampilkan Angka 4 Untuk Setiap Iterasi?

// Ketika JavaScript mengeksekusi kode ini, for loop akan berjalan cepat dan selesai sebelum setTimeout sempat dieksekusi. Ini karena JavaScript memiliki fitur asinkron, yang berarti bahwa setTimeout tidak akan menunda eksekusi kode berikutnya sampai selesai, melainkan akan dilewatkan dan eksekusi kode berikutnya akan dilanjutkan 4.

// Ketika for loop selesai, variabel index akan bernilai 4. Kemudian, setTimeout akan mulai dieksekusi. Karena setTimeout menggunakan konsep closure, ia akan merujuk ke nilai index terakhir, yaitu 4. Oleh karena itu, setiap setTimeout akan mencetak angka 4 4.
 
//B Ada dua cara untuk memperbaiki kode diatas perbaiki kode tersebut sehingga hasilnya
//mencerminkan nilai iterasi yang benar-benar sesuai dengan harapan?.

for (let index = 0; index < 4; index++) {
     setTimeout(()=>{
          console.log(index);
     }, 1000)
 }

 for (var index = 0; index < 4; index++) {
     ((index) => {
         setTimeout(() => {
             console.log(index);
         }, 1000);
     })(index);
 }

//  C Jelaskan mengapa timeout dalam kode tersebut menghasilkan eksekusi asinkron, dan
//  bagaimana konsep event loop di JavaScript berperan dalam hal ini

//  Ketika setTimeout dipanggil, fungsi callback yang dilewatkan ke setTimeout akan ditempatkan dalam antrian event loop dan akan dieksekusi setelah waktu yang ditentukan berakhir. Namun, karena JavaScript adalah single-threaded, kode lain yang ada di stack panggilan akan dieksekusi terlebih dahulu. Dalam kasus ini, for loop akan selesai sebelum setTimeout sempat dieksekusi 4.

// Jadi, meskipun setTimeout memiliki delay 1000 milidetik, delay tersebut tidak mempengaruhi eksekusi for loop. Ini adalah alasan mengapa semua setTimeout mencetak angka 4, karena for loop telah selesai dan variabel index telah menjadi 4 sebelum setTimeout sempat dieksekusi 4.


// 2. Ada yang salah dengan kode dibawah jelaskan apa kesalahanya dan breakdown
// permasalahanya

export function App() {
     const [numberOne, setNumberOne] = useState(0);
     const [numberTwo, setNumberTwo] = useState(0);

     useEffect(()=>{
          setNumberTwo(pre) => pre + 1);
     }, [numberOne]);

     useEffect(()=>{
          setNumberOne(pre) => pre + 1);
     }, [numberTwo]);

     return (
          <div>
               <div>One :{numberOne}</div>
               <div>Two:{numberTwo}</div>
          </div>
     )
}

//keslahan syntax dan cara penempatan parameter di atas berbeda dengan value yang di inginkan dan cara memperbaikinya hanya dengan mengganti ke dua parameter di atas
//hasilperbaikan

export function App() {
     const [numberOne, setNumberOne] = useState(0);
     const [numberTwo, setNumberTwo] = useState(0);
 
     useEffect(() => {
         if (numberOne < 1) {
             setNumberTwo(pre => pre + 1);
         }
     }, [numberOne]);
 
     useEffect(() => {
         if (numberTwo < 1) {
             setNumberOne(pre => pre + 1);
         }
     }, [numberTwo]);
 
     return (
         <div>
             <div>One :{numberOne}</div>
             <div>Two:{numberTwo}</div>
         </div>
     )
 }
// 3. Perhatikan kode dibawah!
// Sekilas tidak ada yang salah dengan kode dibawah ini, tapi bayangkan jika component yang
// digunakan sudah sangat banyak akan sangat sulit melakukan debuging ketika terjadi error.
// Breakdown permasalahan permasalahan yang terjadi dan berika solusinya.
export function App() {
     const [numOfCart, setNumOfCart] = useState(0);
     
     const increment = (e) => {
          setNumOfCart((pre)=> pre + 1);
     };

     return (
          <div>
               <Ui num={numOfCart}>onPress={increment}/>
          </div>
     );

     function Ui({Onpress, num}) {
          return (
               <div>
                     <Total num={num}/>
               <Button onPress={onePress}/>
          </div>
          )
     };

     function Total ({num}) {
          return <div>{num}</div>
     }

     function Button({onPress}) {
          return <button onChange={onPress}>+</button>
     }
}

//cara menyelesainakn permasalahan di atas adalah dengan mebuat component untuk masing masing container dan tidak akan ada penumpukan kode dan mudha di mantain

// 4. Buatlah sebuah web app marketplace sederhana dengan menggunakan stack Golang,
// Nextjs, Firebase dan menerapkan clean code.
//link https://github.com/khairulharu/marketplace