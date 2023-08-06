FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY markdown2bash /markdown2bash
USER 65532:65532

ENTRYPOINT ["/markdown2bash"]
