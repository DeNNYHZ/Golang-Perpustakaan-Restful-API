-- Drop foreign key constraints first to avoid dependency issues
ALTER TABLE IF EXISTS bukus DROP CONSTRAINT IF EXISTS fk_bukus_kategori_jenis;
ALTER TABLE IF EXISTS bukus DROP CONSTRAINT IF EXISTS fk_bukus_penerbit_buku;
ALTER TABLE IF EXISTS bukus DROP CONSTRAINT IF EXISTS fk_bukus_penulis_buku;

ALTER TABLE IF EXISTS dendas DROP CONSTRAINT IF EXISTS fk_anggota_denda;
ALTER TABLE IF EXISTS dendas DROP CONSTRAINT IF EXISTS fk_peminjaman_denda;

ALTER TABLE IF EXISTS detail_pinjams DROP CONSTRAINT IF EXISTS fk_buku_detail_pinjam;
ALTER TABLE IF EXISTS detail_pinjams DROP CONSTRAINT IF EXISTS fk_peminjaman_detail_pinjam;

ALTER TABLE IF EXISTS peminjamen DROP CONSTRAINT IF EXISTS fk_anggota_meminjam;

-- Drop the tables
DROP TABLE IF EXISTS anggota CASCADE;
DROP TABLE IF EXISTS bukus CASCADE;
DROP TABLE IF EXISTS dendas CASCADE;
DROP TABLE IF EXISTS detail_pinjams CASCADE;
DROP TABLE IF EXISTS jenis_bukus CASCADE;
DROP TABLE IF EXISTS pegawais CASCADE;
DROP TABLE IF EXISTS peminjamen CASCADE;
DROP TABLE IF EXISTS penulis_bukus CASCADE;
