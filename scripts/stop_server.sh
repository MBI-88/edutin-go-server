if systemctl is-active --quiet web; then
    echo "El servicio 'web' está en ejecución. Deteniéndolo..."
    sudo systemctl stop web
    if [ $? -eq 0 ]; then
        echo "Servicio 'web' detenido con éxito."
        exit 0
    else
        echo "Error al detener el servicio 'web'."
        exit 1
    fi
else
    echo "El servicio 'web' no está en ejecución."
    exit 0
fi