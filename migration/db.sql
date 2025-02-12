CREATE TABLE kategori (
    ID_Kategori INT PRIMARY KEY AUTO_INCREMENT,
    Nama_Kategori VARCHAR(255) NOT NULL,
    Deskripsi VARCHAR(255)
);

CREATE TABLE produk (
    ID_Produk INT PRIMARY KEY AUTO_INCREMENT,
    Nama_Produk VARCHAR(255) NOT NULL,
    Deskripsi VARCHAR(255),
    Harga DECIMAL(10,2) NOT NULL,
    Stok INT NOT NULL,
    Kategori_ID INT,
    FOREIGN KEY (Kategori_ID) REFERENCES kategori(ID_Kategori)
);

CREATE TABLE pengguna (
    ID_Pengguna INT PRIMARY KEY AUTO_INCREMENT,
    Nama_Pengguna VARCHAR(255) NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Role VARCHAR(50) NOT NULL
);

CREATE TABLE pelanggan (
    ID_Pelanggan INT PRIMARY KEY AUTO_INCREMENT,
    Nama_Pelanggan VARCHAR(255) NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL,
    Telepon VARCHAR(20),
    Alamat VARCHAR(255)
);

CREATE TABLE transaksi (
    ID_Transaksi INT PRIMARY KEY AUTO_INCREMENT,
    Tanggal_Transaksi DATETIME DEFAULT CURRENT_TIMESTAMP,
    Total_Harga DECIMAL(10,2) NOT NULL,
    Pelanggan_ID INT,
    FOREIGN KEY (Pelanggan_ID) REFERENCES pelanggan(ID_Pelanggan)
);

CREATE TABLE detail_transaksi (
    ID_Detail INT PRIMARY KEY AUTO_INCREMENT,
    Transaksi_ID INT,
    Produk_ID INT,
    Jumlah INT NOT NULL,
    Subtotal DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (Transaksi_ID) REFERENCES transaksi(ID_Transaksi),
    FOREIGN KEY (Produk_ID) REFERENCES produk(ID_Produk)
);

CREATE TABLE pembayaran (
    ID_Pembayaran INT PRIMARY KEY AUTO_INCREMENT,
    Transaksi_ID INT,
    Metode_Pembayaran VARCHAR(50) NOT NULL,
    Status_Pembayaran VARCHAR(50) NOT NULL,
    Tanggal_Pembayaran DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Transaksi_ID) REFERENCES transaksi(ID_Transaksi)
);

CREATE TABLE pengiriman (
    ID_Pengiriman INT PRIMARY KEY AUTO_INCREMENT,
    Transaksi_ID INT,
    Alamat_Pengiriman VARCHAR(255) NOT NULL,
    Tanggal_Kirim DATETIME,
    Status_Pengiriman VARCHAR(50) NOT NULL,
    FOREIGN KEY (Transaksi_ID) REFERENCES transaksi(ID_Transaksi)
);

CREATE TABLE riwayat_stok (
    ID_Riwayat INT PRIMARY KEY AUTO_INCREMENT,
    Produk_ID INT,
    Jumlah_Perubahan INT NOT NULL,
    Tanggal_Perubahan DATETIME DEFAULT CURRENT_TIMESTAMP,
    Keterangan VARCHAR(255),
    FOREIGN KEY (Produk_ID) REFERENCES produk(ID_Produk)
);
