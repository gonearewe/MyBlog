$(function () {
    $('#form-register').bootstrapValidator({
      message: 'This value is not valid',
      feedbackIcons: {
        valid: 'glyphicon glyphicon-ok',
        invalid: 'glyphicon glyphicon-remove',
        validating: 'glyphicon glyphicon-refresh'
      },
      live: 'enabled',
      submitButtons: 'button[type="submit"]',
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
    submitHandler: function (form) {
      var urlStr = "/register";
      // alert("urlStr:"+urlStr)
      $(form).ajaxSubmit({
          url:urlStr,
          async:false,
          type:"post",
          dataType:"json",
          cache:false,
          data:{"username":$('#username').val(),"password":$('#password').val(),"repassword":$('#repassword').val()},
          success:function (data,status) {
              alert("data:"+data.message)
              if (data.code == 0){
                  setTimeout(function () {
                      window.location.href="/login"
                  },1000)
              }
          },
          err:function (data,status) {
              alert("err:"+data.message+":"+status)
          }
      })
  }
  });
  
  $('#form-login').bootstrapValidator({
    message: 'This value is not valid',
    submitButtons: 'button[type="submit"]',
    feedbackIcons: {
      valid: 'glyphicon glyphicon-ok',
      invalid: 'glyphicon glyphicon-remove',
      validating: 'glyphicon glyphicon-refresh'
    },
    live: 'enabled',
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
  submitHandler: function (form) {
    var urlStr = "/login";
    alert("urlStr:"+urlStr)
    $(form).ajaxSubmit({
        url:urlStr,
        async:false,
        type:"post",
        dataType:"json",
        cache:false,
        data:{"username":$('#login-username').val(),"password":$('#login-password').val()},
        success:function (data,status) {
            alert("data:"+data.message)
            if (data.code == 0){
                setTimeout(function () {
                    window.location.href="/"
                },1000)
            }
        },
        err:function (data,status) {
            alert("err:"+data.message+":"+status)
        }
    })
}
});

// $('button').click(loginFun)
// function loginFun()
// {
//   $('form').ajaxSubmit({
//     url:'/login',
//     async:false,
//     type:"post",
//     dataType:"json",
//     cache:false,
//     data:{"username":$('#login-username').val(),"password":$('#login-password').val()},
//     success:function (data,status) {
//         alert("data:"+data.message)
//         if (data.code == 0){
//             setTimeout(function () {
//                 window.location.href="/"
//             },1000)
//         }
//     },
//     err:function (data,status) {
//         alert("err:"+data.message+":"+status)
//     }
// })
// }
});

