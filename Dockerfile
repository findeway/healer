FROM alpine
ADD healer-service /healer-service
ENTRYPOINT [ "/healer-service" ]
