all: main #fackmain

# fackmain:
# 	go build fackmain.go

main:
	go build main.go

clean:
	rm ./main