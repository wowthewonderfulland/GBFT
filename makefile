all: main fakemain

fakemain:
	go build fakemain.go

main:
	go build main.go

clean:
	rm ./main ./fakemain