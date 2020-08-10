# Soal 2

Sebuah kumpulan angka disusun secara berurutan dari yang terkecil sampa yang terbesar. Jika ada satu angka yang hilang dalam urutan tersebut, cukup mudah untuk menemukannya. Bila angkanya tidak terlalu banyak, misalnya 10 angka, tanpa bantuan alat, manusia masih bisa mencarinya. Misalnya pada urutan angka berikut ini.

`14 15 16 18 19 20`

Dapat dilihat dengan mudah angka yang hilang adalah 17. Sekarang bagaimana jika spasi dalam susunan angka tersebut kita hilangkan? Susunan angka menjadi seperti berikut ini.

`141516181920`

Dengan hanya memberi informasi bahwa angka tersebut berurutan, dan ada satu angka yang hilang dalam urutan tersebut, menemukan angka yang hilang tidak lagi semudah sebelumnya. Ada usaha ekstra yang diperlukan.

### Data soal

Setiap baris berisi angka berurutan dari kecil ke besar. Angka minimal adalah 1, dan angka maksimal adalah 10 pangkat 6. Tidak ada pemisah antara angka dalam urutan tersebut. Panjang minimal satu baris soal adalah 3 angka, maksimal 1000 angka. Angka yang hilang tidak berada di awal atau akhir, melainkan berada di antara awal dan akhir. Program dibuat untuk menemukan angka yang hilang.

### Contoh

```
Input                                   Output

23242526272830                          29
101102103104105106107108109111112113    110
12346789                                5
79101112                                8
7891012                                 11
9799100101102                           98                              
100001100002100003100004100006          100005

```