//Prompt is the JavaScript module for all alerts, notifications and custom popup dialogs
function Prompt() {

    let prompt_alert = function (icon, c) {
        const {
            title = '',
            text = '',
            footer = ''
        } = c
        Swal.fire({
            icon: icon,
            title: title,
            text: text,
            footer: footer
        })
    }

    let success = prompt_alert('success', c)

    let error = prompt_alert('error', c)

    let warning = prompt_alert('warning', c)

    async function custom(c) {
        const {
            icon = '',
            title = '',
            html = '',
            showConfirmButton = true
        } = c
        const {value: result} = await Swal.fire({
            icon: icon,
            title: title,
            html: html,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton: showConfirmButton,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen();
                }
            },
        })
        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result !== "") {
                    if (c.callback !== undefined) {
                        c.callback(result);
                    } else {
                        c.callback(false);
                    }
                }
            } else {
                c.callback(false);
            }
        }
    }

    return {
        success: success,
        warning: warning,
        error: error,
        custom: custom
    }
}