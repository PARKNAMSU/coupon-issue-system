USE mysql;

CREATE TABLE `campagin` ( -- 캠페인 테이블
    `campagin_id` INT AUTO_INCREMENT PRIMARY KEY COMMENT "캠페인 아이디",
    `name` VARCHAR(50) DEFAULT NULL COMMENT "캠페인 이름",
    `total_coupons` INT NOT NULL DEFAULT 0 COMMENT "최대 발급 가능한 쿠폰 개수",
    `status` INT NOT NULL DEFAULT 0 COMMENT "캠페인 상태 0: 닫힘, 1: 열림",
    `avaliable_at` TIMESTAMP NOT NULL COMMENT "캠페인 오픈 시간",
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT "캠페인 생성 시간"
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `coupon` ( -- 쿠폰 등록 테이블
    `campagin_id` INT AUTO_INCREMENT PRIMARY KEY COMMENT "연결된 캠페인 아이디",
    `coupon_code` VARCHAR(50) NOT NULL COMMENT "쿠폰 코드",
    `receive_account` VARCHAR(50) NOT NULL COMMENT "전달받을 계정",
    `receive_method` VARCHAR(10) NOT NULL COMMENT "전달받을 방법 ex) email, phone, account ...",
    `is_used` TINYINT(1) NOT NULL DEFAULT 0 COMMENT "쿠폰 사용 여부", 
    `is_received` TINYINT(1) NOT NULL DEFAULT 0 COMMENT "쿠폰 전달 여부",
    `used_at` TIMESTAMP NOT NULL COMMENT "쿠폰 사용 시간",
    `received_at` TIMESTAMP NOT NULL COMMENT "쿠폰 전달 시간",
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT "쿠폰 사용 시간",
    PRIMARY KEY(`campagin_id`,`coupon_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;