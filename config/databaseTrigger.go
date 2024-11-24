package config

import "gorm.io/gorm"

// membuat trigger ketika produk per perusahaan ditambahkan sehingga membuat parameter stok produk sebanyak instansi yang ada di perusahaan tersebut
func CreateTriggerAddInstansiStokProduk(db *gorm.DB) error {
	triggerSQL := `
	CREATE OR REPLACE FUNCTION add_instansistokproduks_after_insertProduk() 
	RETURNS TRIGGER AS $$
	BEGIN
    INSERT INTO instansistokproduks (id, instansi_id, produk_id, stok, created_at, updated_at)
    SELECT uuid_generate_v4(), i.id, NEW.id, 0, NEW.created_at, NEW.updated_at
    FROM instansis i
    WHERE i.perusahaan_id = NEW.perusahaan_id;
    RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	DROP TRIGGER IF EXISTS add_instansistokproduks_after_insertProduk ON products;

	CREATE TRIGGER add_instansistokproduks_after_insertProduk
	AFTER INSERT ON products
	FOR EACH ROW
	EXECUTE FUNCTION add_instansistokproduks_after_insertProduk();
	`
	// Eksekusi perintah SQL untuk membuat trigger
	return db.Exec(triggerSQL).Error
}

// Membuat trigger barang masuk dan menambah kumulatif stok di tabel produk setelah migrasi selesai
func CreateTriggerBrMasukTambhStokSQL(db *gorm.DB) error {
	triggerSQL := `
	CREATE OR REPLACE FUNCTION add_kumulatifstok_after_insertBarangMasuk() 
	RETURNS TRIGGER AS $$
	BEGIN
		UPDATE products SET kumulatif_stok = kumulatif_stok + NEW.qty WHERE produk_id = NEW.produk_id;
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	DROP TRIGGER IF EXISTS add_kumulatifstok_after_insertBarangMasuk ON barangmasuks;

	CREATE TRIGGER add_kumulatifstok_after_insertBarangMasuk
	AFTER INSERT ON barangmasuks
	FOR EACH ROW
	EXECUTE FUNCTION add_kumulatifstok_after_insertBarangMasuk();
	`
	// Eksekusi perintah SQL untuk membuat trigger
	return db.Exec(triggerSQL).Error
}

// Membuat trigger barang masuk dan menambah kumulatif stok di tabel produk setelah migrasi selesai
func CreateTriggerBrMasukKurangiStokSQL(db *gorm.DB) error {

	triggerSQL := `
	CREATE OR REPLACE FUNCTION kurangi_kumulatifstok_after_deleteBarangMasuk() 
	RETURNS TRIGGER AS $$
	BEGIN
	    UPDATE products SET kumulatif_stok = kumulatif_stok - OLD.qty WHERE produk_id = OLD.produk_id;
	    RETURN OLD;
	END;
	$$ LANGUAGE plpgsql;

	DROP TRIGGER IF EXISTS kurangi_kumulatifstok_after_deleteBarangMasuk ON barangmasuks;

	CREATE TRIGGER kurangi_kumulatifstok_after_deleteBarangMasuk
	AFTER DELETE ON barangmasuks
	FOR EACH ROW
	EXECUTE FUNCTION kurangi_kumulatifstok_after_deleteBarangMasuk();
	`
	// Eksekusi perintah SQL untuk membuat trigger
	return db.Exec(triggerSQL).Error
}

// Membuat trigger barang distribusi unit dan menambah stok di tabel instansistokproduks
func CreateTriggerBrDistribusiUnitTambahStokSQL(db *gorm.DB) error {

	triggerSQL := `
	CREATE OR REPLACE FUNCTION tambah_stok_produk_instansi_after_insertBarangdistribusikantors() 
	RETURNS TRIGGER AS $$
	BEGIN
	    UPDATE instansistokproduks SET stok = stok - NEW.qty WHERE produk_id = NEW.produk_id;
	    RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	DROP TRIGGER IF EXISTS tambah_stok_produk_instansi_after_insertBarangdistribusikantors ON barangdistribusikantors;

	CREATE TRIGGER tambah_stok_produk_instansi_after_insertBarangdistribusikantors
	AFTER INSERT ON barangdistribusikantors
	FOR EACH ROW
	EXECUTE FUNCTION tambah_stok_produk_instansi_after_insertBarangdistribusikantors();
	`
	// Eksekusi perintah SQL untuk membuat trigger
	return db.Exec(triggerSQL).Error
}

func CreateTriggerBrDistribusiUnitKurangiStokSQL(db *gorm.DB) error {
	triggerSQL := `
	CREATE OR REPLACE FUNCTION kurangi_stok_produk_instansi_after_deleteBarangdistribusikantors() 
	RETURNS TRIGGER AS $$
	BEGIN
    UPDATE instansistokproduks SET stok = stok - OLD.qty WHERE produk_id = OLD.produk_id;
    RETURN OLD;
	END;
	$$ LANGUAGE plpgsql;

	DROP TRIGGER IF EXISTS kurangi_stok_produk_instansi_after_deleteBarangdistribusikantors ON barangdistribusikantors;

	CREATE TRIGGER kurangi_stok_produk_instansi_after_deleteBarangdistribusikantors
	AFTER DELETE ON barangdistribusikantors
	FOR EACH ROW
	EXECUTE FUNCTION kurangi_stok_produk_instansi_after_deleteBarangdistribusikantors();
	`
	// Eksekusi perintah SQL untuk membuat trigger
	return db.Exec(triggerSQL).Error
}
func CreateTriggerAddActivityDetail(db *gorm.DB) error {
	triggerSQL := `
	CREATE OR REPLACE FUNCTION insert_empactivity_details()
	RETURNS TRIGGER AS $$
	BEGIN
    INSERT INTO emp_activity_detail_mandors (activity_id, employee_id)
    VALUES (NEW.id, NEW.mandorsatu_id);

    INSERT INTO emp_activity_detail_keranis (activity_id, employee_id)
    VALUES (NEW.id, NEW.kerani_id);

    INSERT INTO emp_activity_detail_extras (activity_id, employee_id)
    VALUES (NEW.id, NEW.extra_id);

    RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;


	DROP TRIGGER IF EXISTS insert_empactivity_details ON emp_activities;

	CREATE TRIGGER insert_empactivity_details
	AFTER INSERT ON emp_activities
	FOR EACH ROW
	EXECUTE FUNCTION insert_empactivity_details();
	`
	// Eksekusi perintah SQL untuk membuat trigger
	return db.Exec(triggerSQL).Error
}
