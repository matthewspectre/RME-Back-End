CREATE TABLE rencana_rawat (
    id_rencana_rawat INT NOT NULL AUTO_INCREMENT,
    id_pasien INT NOT NULL,
    id_dokter INT NOT NULL,
    rencana_rawat TEXT NOT NULL,
    date_make DATETIME NOT NULL,
    date_update DATETIME NOT NULL,
    visible TINYINT NOT NULL DEFAULT 1,
    PRIMARY KEY (id_rencana_rawat),
    KEY idx_rencana_rawat_pasien (id_pasien),
    KEY idx_rencana_rawat_dokter (id_dokter),
    KEY idx_rencana_rawat_visible_date (visible, date_make)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
