curl -X POST --data '{
 "file": "local:/home/paul/IdeaProjects/lfs-monthly/target/scala-2.11/lfs-monthly-assembly-1.0.jar",
 "name": "My Test Job",
 "className": "uk.gov.ons.lfs.LFSMonthly"}'  -H "Content-Type: application/json" http://localhost:8998/batches
