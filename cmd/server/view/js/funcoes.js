function enviaValores() {

    var integral = {
        corpo: '',
        parametro: '',
        a: 0,
        b: 0
    };

    calculo = $('input:radio[name=calculo]:checked').val()

    console.log(calculo);

    integral.corpo = $('#corpo').val();
    integral.parametro = $('#parametro').val();
    integral.a = $('#a').val();
    integral.b = $('#b').val();
    erro = $('#erro').val();

    axios.post('/' + calculo + '/' + erro, integral)
        .then(function (response) {
            console.log(response.data.result);
            $('#resultado').text("Resultado: " + response.data.result)
            $('#resultado').removeClass("alert-danger")
            $('#resultado').addClass("alert-success")
            $('#resultado').removeAttr('hidden')
        })
        .catch(function (error) {
            $('#resultado').text(error)
            $('#resultado').removeClass("alert-success")
            $('#resultado').addClass("alert-danger")
            $('#resultado').removeAttr('hidden')
            console.log(error);
        });
}

$(window).resize(function () {
    if ($(window).width() <= 800) {
        document.getElementById('menu').className = 'btn-group-vertical item centraliza d-flex justify-content-center';
    } else if ($(window).width() > 550) {
        document.getElementById('menu').className = 'btn-group item centraliza d-flex justify-content-center';
    }
});



