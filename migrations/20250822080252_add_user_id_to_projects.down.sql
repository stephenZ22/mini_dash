-- 先删掉外键约束
ALTER TABLE projects DROP CONSTRAINT IF EXISTS fk_projects_user;

-- 再删索引
DROP INDEX IF EXISTS idx_projects_user_id;

-- 最后删掉列
ALTER TABLE projects DROP COLUMN IF EXISTS user_id;
