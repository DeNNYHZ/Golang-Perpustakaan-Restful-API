-- Table for members (anggota)
CREATE TABLE anggota (
                         id_anggota VARCHAR(26) PRIMARY KEY,
                         username VARCHAR(255) UNIQUE NOT NULL,
                         password VARCHAR(100) NOT NULL,
                         nama VARCHAR(255) NOT NULL,
                         sex VARCHAR(255) NOT NULL,
                         telp VARCHAR(255) NOT NULL,
                         alamat TEXT NOT NULL,
                         email VARCHAR(255) UNIQUE NOT NULL,
                         deskripsi TEXT NOT NULL,
                         created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for books (bukus)
CREATE TABLE bukus (
                       id_buku VARCHAR(26) PRIMARY KEY,
                       isbn VARCHAR(255) UNIQUE NOT NULL,
                       id_kategori VARCHAR(26) NOT NULL,
                       judul_buku VARCHAR(255) NOT NULL,
                       id_penulisbuku VARCHAR(26),
                       id_penerbitbuku VARCHAR(26),
                       tahun_terbit VARCHAR(255),
                       stok_buku INT,
                       rak_buku VARCHAR(255),
                       deskripsi_buku TEXT,
                       gambar_buku VARCHAR(255),
                       kondisi_buku VARCHAR(255),
                       created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for fines (dendas)
CREATE TABLE dendas (
                        id_denda VARCHAR(26) PRIMARY KEY,
                        jumlah_denda INT NOT NULL,
                        tglpinjam TIMESTAMP(3),
                        tglhrskembali TIMESTAMP(3),
                        tglkembali TIMESTAMP(3),
                        id_peminjaman VARCHAR(26) NOT NULL,
                        id_anggota VARCHAR(26) NOT NULL,
                        created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for borrowing details (detail_pinjams)
CREATE TABLE detail_pinjams (
                                id_detailpinjam VARCHAR(26) PRIMARY KEY,
                                id_buku VARCHAR(26),
                                id_peminjaman VARCHAR(26),
                                kondisi VARCHAR(255) NOT NULL,
                                created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for book categories (jenis_bukus)
CREATE TABLE jenis_bukus (
                             id_jenis VARCHAR(26) PRIMARY KEY,
                             jenis_buku VARCHAR(255) UNIQUE NOT NULL,
                             deskripsi TEXT,
                             updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for employees (pegawais)
CREATE TABLE pegawais (
                          id_pegawai VARCHAR(26) PRIMARY KEY,
                          username VARCHAR(255) UNIQUE NOT NULL,
                          password VARCHAR(100) NOT NULL,
                          created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for borrowings (peminjaman)
CREATE TABLE peminjaman (
                            id_peminjaman VARCHAR(26) PRIMARY KEY,
                            id_anggota VARCHAR(26),
                            tgl_pinjam TIMESTAMP(3),
                            tgl_hrs_kembali TIMESTAMP(3),
                            jaminan VARCHAR(255) NOT NULL,
                            created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for book publishers (penerbit_bukus)
CREATE TABLE penerbit_bukus (
                                id_penerbit VARCHAR(26) PRIMARY KEY,
                                penerbit_buku VARCHAR(255) UNIQUE NOT NULL,
                                alamat_penerbit VARCHAR(255),
                                telp_penerbit VARCHAR(255),
                                email_penerbit VARCHAR(255) UNIQUE,
                                deskripsi TEXT,
                                updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Table for book authors (penulis_bukus)
CREATE TABLE penulis_bukus (
                               id_penulis VARCHAR(26) PRIMARY KEY,
                               penulis_buku VARCHAR(255) UNIQUE NOT NULL,
                               alamat_penulis VARCHAR(255),
                               email_penulis VARCHAR(255) UNIQUE,
                               deskripsi TEXT,
                               updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP
);

-- Foreign key constraints

-- Constraints for table bukus
ALTER TABLE bukus
    ADD CONSTRAINT fk_bukus_kategori_jenis FOREIGN KEY (id_kategori) REFERENCES jenis_bukus (id_jenis),
  ADD CONSTRAINT fk_bukus_penerbit_buku FOREIGN KEY (id_penerbitbuku) REFERENCES penerbit_bukus (id_penerbit),
  ADD CONSTRAINT fk_bukus_penulis_buku FOREIGN KEY (id_penulisbuku) REFERENCES penulis_bukus (id_penulis);

-- Constraints for table dendas
ALTER TABLE dendas
    ADD CONSTRAINT fk_anggota_denda FOREIGN KEY (id_anggota) REFERENCES anggota (id_anggota),
  ADD CONSTRAINT fk_peminjaman_denda FOREIGN KEY (id_peminjaman) REFERENCES peminjamen (id_peminjaman);

-- Constraints for table detail_pinjams
ALTER TABLE detail_pinjams
    ADD CONSTRAINT fk_buku_detail_pinjam FOREIGN KEY (id_buku) REFERENCES bukus (id_buku),
  ADD CONSTRAINT fk_peminjaman_detail_pinjam FOREIGN KEY (id_peminjaman) REFERENCES peminjamen (id_peminjaman);

-- Constraints for table peminjamen
ALTER TABLE peminjamen
    ADD CONSTRAINT fk_anggota_meminjam FOREIGN KEY (id_anggota) REFERENCES anggota (id_anggota);
