ALTER TABLE default.test_table ADD COLUMN NewField UInt32;

/* rollback
ALTER TABLE default.test_table DROP COLUMN NewField;
*/
