-- アップロードされた画像を保存するテーブル
-- 画像の内容は image テーブルに保存される
CREATE TABLE IF NOT EXISTS contents (
    id INTEGER NOT NULL AUTO_INCREMENT,

    name VARCHAR(255) NOT NULL,
    extension VARCHAR(255),

    image_id INTEGER NOT NULL,

    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
) CHARSET=utf8;

-- 画像の内容を保存するテーブル
-- 画像の内容は content テーブルに保存される
CREATE TABLE IF NOT EXISTS images (
    id INTEGER NOT NULL AUTO_INCREMENT,

    content LONGTEXT NOT NULL,
    sha256_digest VARCHAR(64) NOT NULL,

    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    UNIQUE KEY sha256_digest (sha256_digest)
) CHARSET=utf8;