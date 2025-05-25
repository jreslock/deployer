# syntax=docker/dockerfile:1

FROM public.ecr.aws/lambda/provided:al2023

ARG TARGETARCH

COPY artifacts/deployer-linux-${TARGETARCH} ./deployer

ENTRYPOINT [ "./deployer" ]
