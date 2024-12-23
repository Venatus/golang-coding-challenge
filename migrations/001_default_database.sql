CREATE TABLE IF NOT EXISTS default.test_table
(
    `TenantId` UInt8,
    `AccountId` UInt16,
    `SiteId` UInt32,
    `Time` DateTime,
	`Created` DateTime DEFAULT NOW(),
    `Url` String,
    `Ip` Nullable(IPv4),
    `Ipv6` Nullable(IPv6),
    `Geo` LowCardinality(String),
    `UserAgent` String,
    `Device` LowCardinality(String)
)
ENGINE = MergeTree
PRIMARY KEY (toStartOfHour(`Time`), TenantId, AccountId, SiteId)
ORDER BY (toStartOfHour(`Time`), TenantId, AccountId, SiteId)
SETTINGS index_granularity = 8192;


/* rollback
DROP TABLE IF EXISTS default.test_table
*/
