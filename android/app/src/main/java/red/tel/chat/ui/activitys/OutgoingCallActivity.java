package red.tel.chat.ui.activitys;import android.os.Bundle;import android.support.annotation.Nullable;import android.view.View;import android.widget.Button;import android.widget.TextView;import java.nio.ByteBuffer;import java.nio.ShortBuffer;import red.tel.chat.EventBus;import red.tel.chat.R;import red.tel.chat.VoipBackend;import red.tel.chat.generated_protobuf.Voip;import red.tel.chat.network.NetworkCall;import red.tel.chat.network.NetworkCallInfo;import red.tel.chat.network.NetworkCallProposalInfo;import red.tel.chat.network.OutgoingCallProposalController;import red.tel.chat.ui.fragments.ItemDetailFragment;import static red.tel.chat.generated_protobuf.Voip.Which.CALL_DECLINE;import static red.tel.chat.generated_protobuf.Voip.Which.CALL_STOP;import static red.tel.chat.ui.activitys.IncomingCallActivity.TYPE_CALL;public class OutgoingCallActivity extends BaseCall implements View.OnClickListener {    private boolean isVideo = false;    private TextView from;    private Button btnCancel;    private NetworkCallProposalInfo callInfo;    private NetworkCall networkCall;    @Override    protected void onCreate(@Nullable Bundle savedInstanceState) {        super.onCreate(savedInstanceState);        setContentView(R.layout.activity_outgoing_call);        from = findViewById(R.id.from);        cameraView = findViewById(R.id.camera);        btnCancel = findViewById(R.id.cancel);        btnCancel.setOnClickListener(this);        networkCall = new NetworkCall() {            @Override            public NetworkCallInfo getNetworkCallInfo() {                return super.getNetworkCallInfo();            }        };        Bundle bundle = getIntent().getExtras();        if (bundle != null) {            isVideo = bundle.getBoolean(TYPE_CALL);            String whom = bundle.getString(ItemDetailFragment.ARG_ITEM_ID);            if (isVideo) {                callInfo =  networkCall.callVideoAsync(whom);            } else {                callInfo = networkCall.callAudioAsync(whom);            }            from.setText(whom != null ? "Call to: " + whom : "");        }        cameraView.setVisibility(isVideo ? View.VISIBLE : View.GONE);    }    @Override    protected void onSubscribeEvent(Object object) {        super.onSubscribeEvent(object);        if (object == EventBus.Event.STOP_CALL ||                (object instanceof Voip && ((Voip) object).which == CALL_DECLINE)                || (object instanceof Voip && ((Voip) object).which == CALL_STOP)) {            finish();        }        if (object.equals("ACCEPT")) {            requestPermissions();        }    }    @Override    protected void onCallBackRecord(ByteBuffer buffer, ShortBuffer[] samples) {        VoipBackend.getInstance().sendDataToServerWhenAccept(buffer, samples);    }    @Override    protected boolean isVideo() {        return isVideo;    }    @Override    public void onClick(View view) {        switch (view.getId()) {            case R.id.cancel:                if (callInfo != null) {                    OutgoingCallProposalController.getInstance().stop(callInfo);                    finish();                }                break;        }    }    @Override    protected void onDestroy() {        super.onDestroy();        callInfo = null;    }}