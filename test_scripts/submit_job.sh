
now=$(date +%s)
job="Test Livy $now"

jar="local:/home/paul/IdeaProjects/lfs-monthly/target/scala-2.11/lfs-monthly-assembly-1.0.jar"

str=$(printf '{"job":"%s","jar":"%s"}\n' "$job" "$jar")

curl -X POST --data "$str" -H "Content-Type: application/json" localhost:4000/submit
