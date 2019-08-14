$(function () {
    $('form').bootstrapValidator({
      message: 'This value is not valid',
      feedbackIcons: {
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
      },
      fields: {
        username: {
          message: '用户名验证失败',
          validators: {
            notEmpty: {
              message: '用户名不能为空'
            },
            stringLength: {
              min: 4,
              max: 18,
              message: '用户名长度必须在4到18位之间'
            },
            regexp: {
              regexp: /^[a-zA-Z0-9_]+$/,
              message: '用户名只能包含大写、小写、数字和下划线'
            }
          }
        },
        password: {
          validators: {
            notEmpty: {
              message: '密码不能为空'
            },
            stringLength: {
                min: 8,
                max: 24,
                message: '密码长度必须在8到24位之间'
              },
            regexp: {
                regexp: /^[a-zA-Z0-9_]+$/,
                message: '密码只能包含大写、小写、数字和下划线'
            }
        }
      }
    },
    submitHandler: function (validator, form, submitButton) {
      alert("submit");
    }
  });
});