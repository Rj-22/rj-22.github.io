(function ($) {
  "use strict";

  function use_recaptcha() {
    const site_key = '6LfvyfoUAAAAADV3kOFzYjXDTJtkANsp_cB3Allo';

    grecaptcha.ready(function () {
      grecaptcha.execute(site_key, {action: 'homepage'}).then(function (token) {
        $("form.use_recaptcha").each(function () {
          $(this).append('<input type="hidden" name="g-recaptcha-response" value="' + token + '" > ');
        });
      });
    });
  }

  if ($("form.use_recaptcha").length > 0) {
    use_recaptcha();
  }

  $("form.mf_form_validate").each(function () {
    $(this).validate({
      rules: {
        name: {
          required: true,
          minlength: 2
        },
        email: {
          required: true,
          email: true
        },
        phone: {
          required: true,
          number: true
        },
        password: "required",
        repeat_password: {
          equalTo: "#password_field"
        }
      }
    });
  });

  $("form.ajax_submit").on('submit', function () {
    let has_errors = false,
      form = $(this),
      form_fields = form.find('input'),
      form_message = form.find('textarea'),
      server_result_display = form.find('.server_response');

    const poster_title = form.find('[name=poster_title]').val(),
      email = form.find('[name=email]').val(),
      phone = form.find('[name=phone]').val(),
      subject = form.find('[name=subject]').val(),
      message = form.find('[name=message]').val();

    form_fields.each(function () {
      if ($(this).hasClass('error')) {
        has_errors = true;
      }
    });

    if (form_message.length > 0) {
      if (form_message.hasClass('error')) {
        has_errors = true;
      }
    }

    const data_json = JSON.stringify({
      'poster_title': poster_title,
      'email': email,
      'phone': phone,
      'subject': subject,
      'message': message,
    });

    if (!has_errors) {
      $.ajax({
        type: "POST",
        url: form.attr('url'),
        data: data_json,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data) {
          const response = data;
          if (response.status === 'error') {
            server_result_display.empty().html('<div class="mf_alert danger">' + response.errors + '</div>');
          } else if (response.status === 'success') {
            server_result_display.empty().html('<div class="mf_alert success">' + response.message + '</div>');
            setTimeout(function () {
              $('form.ajax_submit .mf_alert').fadeOut(500);
              window.location.replace('https://idpay.ir/rojdax');
            }, 2000);
            form.trigger("reset");
          }
        },
        error: function () {
          server_result_display.empty().html('<div class="mf_alert danger">Server error! Please try again...</div>');
        }
      });
    }

    return false;
  });
})(jQuery);
