# golang_leetcode

## GPG Key Generation

```sh
gpg --full-generate-key
gpg --list-secret-keys --keyid-format LONG
gpg --armor --export `key-ID`
gpg --export-secret-keys `key-ID` > /Users/timhsu/dev_projects/key/private.key
gpg --armor --export-secret-key timhsu.2030@hotmail.com > /Users/timhsu/dev_projects/key/private.gpg
gpg --delete-secret-key `key-ID`
gpg --delete-key `key-ID`
```

## Import GPG

```sh
apt install -y gawk
rm -rf ./git-secret
git clone https://github.com/sobolevn/git-secret.git
cd git-secret
make build
PREFIX="/usr/local" 
make install
cd ..
rm -rf ./git-secret

gpg --import /trade_env/private.gpg
gpg --edit-key timhsu.2030@hotmail.com
```

## Git-Secret

```sh
git secret init
git secret tell timhsu.2030@hotmail.com
git secret add env.sh
git secret hide
git secret reveal
```

```py
def print_to_stdout(*a):
    print(*a, file=sys.stdout)
    sys.stdout.flush()

i = int()
while i <= 10:
    print_to_stdout(i)
    # print(i)
    i += 1
    time.sleep(0.3)
    if i == 10:
        i = 1
```

```go
qb, _ := orm.NewQueryBuilder("mysql")
qb.Select("*").
    From("tick_entire").
    Where("timestamp > ?").And("timestamp < ?").
    OrderBy("timestamp").Asc()
sql := qb.String()
o.Raw(sql, startTime, endTime).QueryRows(&result)
```

```go
_, err = db.Exec("DROP TABLE IF EXISTS " + dbName + ".basic_stock;")
if err != nil {
 panic(err)
}

sort.Slice(res, func(i, j int) bool {
    return res[i].TS < res[j].TS
})
```

## proto

```sh
git reset -- hard
git clean -fxd
go mod tidy
echo '#!/bin/bash
protoc --go_out=./ ./pb/commonpb/enum.proto

protoc --go_out=./ ./pb/gatewaypb/data.proto
protoc --go_out=./ ./pb/gatewaypb/event.proto
protoc --go_out=./ ./pb/gatewaypb/query.proto

protoc --go_out=./ ./pb/internalpb/control.proto
protoc --go_out=./ ./pb/internalpb/device_properties.proto
protoc --go_out=./ ./pb/internalpb/network.proto
protoc --go_out=./ ./pb/internalpb/preference.proto

mkdir ./pkg/core/pb/gatewaypb

mv ./moxa/mxview/pkg/core/pb/commonpb/enum.pb.go ./pkg/core/pb/commonpb/enum.pb.go

mv ./moxa/mxview/pkg/core/pb/gatewaypb/data.pb.go ./pkg/core/pb/gatewaypb/data.pb.go
mv ./moxa/mxview/pkg/core/pb/gatewaypb/event.pb.go ./pkg/core/pb/gatewaypb/event.pb.go
mv ./moxa/mxview/pkg/core/pb/gatewaypb/query.pb.go ./pkg/core/pb/gatewaypb/query.pb.go

mv ./moxa/mxview/pkg/core/pb/internalpb/control.pb.go ./pkg/core/pb/internalpb/control.pb.go
mv ./moxa/mxview/pkg/core/pb/internalpb/device_properties.pb.go ./pkg/core/pb/internalpb/device_properties.pb.go
mv ./moxa/mxview/pkg/core/pb/internalpb/network.pb.go ./pkg/core/pb/internalpb/network.pb.go
mv ./moxa/mxview/pkg/core/pb/internalpb/preference.pb.go ./pkg/core/pb/internalpb/preference.pb.go ' > proto.sh
chmod +x ./proto.sh
./proto.sh
```
