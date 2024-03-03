CREATE TABLE IF NOT EXISTS "default"."logs"
(
    "Id" Int,
    "ProjectId" Int,
    "Name" String,
    "Description" String,
    "Priority" Int,
    "Removed" Bool,
    "EventTime" DateTime DEFAULT now()
)
ENGINE MergeTree()
ORDER BY Id


