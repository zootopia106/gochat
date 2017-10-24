package red.tel.chat.office365;


public class Constants {
    public static final String AUTHORITY_URL = "https://login.microsoftonline.com/common";
    public static final String AUTHORIZATION_ENDPOINT = "/oauth2/v2.0/authorize";
    public static final String TOKEN_ENDPOINT = "/oauth2/v2.0/token";
    // Update these two constants with the values for your application:
    public static String CLIENT_ID = "044e6315-8adc-4dce-9e8e-d9c4d8fef806";
    public static final String REDIRECT_URI = "https://login.microsoftonline.com/common/oauth2/nativeclient";
    public static final String SCOPES = "openid profile mail.send contacts.read user.read";

    public static final String API_ENDPOINT = "https://graph.microsoft.com/beta/";
    public static final int TYPE_LOGIN_MS = 2;
    public static final int TYPE_LOGIN_NORMAL = 1;
    public interface ACTION {
        String MAIN_ACTION = "action.main";
        String STARTFOREGROUND_ACTION = "action.startforeground";
        String STOPFOREGROUND_ACTION = "action.stopforeground";
    }

    public interface NOTIFICATION_ID {
        int FOREGROUND_SERVICE = 101;
    }
}