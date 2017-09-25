package red.tel.chat.ui.activitys;import android.hardware.Camera;import android.media.AudioTrack;import android.os.Bundle;import android.support.annotation.Nullable;import android.util.Log;import android.view.TextureView;import android.view.View;import android.widget.Button;import android.widget.TextView;import java.nio.ByteBuffer;import java.nio.ShortBuffer;import java.util.Arrays;import java.util.concurrent.TimeUnit;import io.reactivex.Observable;import io.reactivex.Observer;import io.reactivex.Scheduler;import io.reactivex.disposables.Disposable;import io.reactivex.schedulers.Schedulers;import red.tel.chat.EventBus;import red.tel.chat.R;import red.tel.chat.VoipBackend;import red.tel.chat.generated_protobuf.Image;import red.tel.chat.generated_protobuf.Voip;import red.tel.chat.io.IO;import red.tel.chat.network.NetworkCall;import red.tel.chat.network.NetworkCallInfo;import red.tel.chat.network.NetworkCallProposalInfo;import red.tel.chat.network.OutgoingCallProposalController;import red.tel.chat.ui.fragments.ItemDetailFragment;import static red.tel.chat.generated_protobuf.Voip.Which.CALL_DECLINE;import static red.tel.chat.generated_protobuf.Voip.Which.CALL_STOP;import static red.tel.chat.ui.activitys.IncomingCallActivity.TYPE_CALL;public class OutgoingCallActivity extends BaseCall implements View.OnClickListener, IO.IODataProtocol {    private static final String TAG = OutgoingCallActivity.class.getSimpleName();    private boolean isVideo = false;    private TextView from;    private Button btnCancel;    private View viewRoot;    private NetworkCallProposalInfo callInfo;    private NetworkCall networkCall;    private IO.IOID ioid;    @Override    protected void onCreate(@Nullable Bundle savedInstanceState) {        super.onCreate(savedInstanceState);        setContentView(R.layout.activity_outgoing_call);        from = findViewById(R.id.from);        mCameraTextureView = findViewById(R.id.camera);        btnCancel = findViewById(R.id.cancel);        viewRoot = findViewById(R.id.root);        btnCancel.setOnClickListener(this);        networkCall = new NetworkCall() {            @Override            public NetworkCallInfo getNetworkCallInfo() {                return super.getNetworkCallInfo();            }        };        Bundle bundle = getIntent().getExtras();        if (bundle != null) {            isVideo = bundle.getBoolean(TYPE_CALL);            String whom = bundle.getString(ItemDetailFragment.ARG_ITEM_ID);            if (isVideo) {                callInfo =  networkCall.callVideoAsync(whom);            } else {                callInfo = networkCall.callAudioAsync(whom);            }            from.setText(whom != null ? "Call to: " + whom : "");        }        mCameraTextureView.setVisibility(isVideo ? View.VISIBLE : View.GONE);        VoipBackend.getInstance().setIoDataProtocol(this);        ioid = new IO.IOID(callInfo.from, callInfo.to, callInfo.getId(), callInfo.getId());    }    @Override    protected void onSubscribeEvent(Object object) {        super.onSubscribeEvent(object);        if (object == EventBus.Event.STOP_CALL ||                (object instanceof Voip && ((Voip) object).which == CALL_DECLINE)                || (object instanceof Voip && ((Voip) object).which == CALL_STOP)) {            finish();        }    }    @Override    protected void onCallBackRecord(ByteBuffer buffer, ShortBuffer[] samples, byte[] data) {        VoipBackend.getInstance().sendDataAudioToServerWhenAccept(data, ioid);    }    @Override    protected void onCallVideoData(byte[] data, Camera.Size size) {        VoipBackend.getInstance().sendDataVideoToServerWhenAccept(data, size, ioid);    }    @Override    protected boolean isVideo() {        return isVideo;    }    @Override    public void onClick(View view) {        switch (view.getId()) {            case R.id.cancel:                if (callInfo != null) {                    OutgoingCallProposalController.getInstance().stop(callInfo);                    finish();                }                break;        }    }    @Override    protected void onDestroy() {        super.onDestroy();        callInfo = null;    }    @Override    public void startOut() {        requestPermissions();    }    @Override    public void processAudio(byte[] data) {        Log.d(TAG, "processAudio: " + Arrays.toString(data));        if (audioTrack != null) {            if (AudioTrack.PLAYSTATE_PLAYING != audioTrack.getPlayState()) {                audioTrack.play();                Log.d(TAG, "Play audio: ");            }            int size = audioTrack.write(data, 0, data.length);            if (data.length != size) {                Log.i(TAG, "Failed to send all data to audio output, expected size: " +                        data.length + ", actual size: " + size);            }        }    }    @Override    public void processVideo(Image data) {        Log.d(TAG, "processVideo: "+ Arrays.toString(data.data.toByteArray()) + " " + data.height + " " + data.width);    }    @Override    public void setSurfaceTextureListener(TextureView.SurfaceTextureListener listener) {        if (mCameraTextureView != null) {            mCameraTextureView.setSurfaceTextureListener(listener);        }    }    @Override    public TextureView getTextureView() {        return mCameraTextureView;    }    @Override    public int getUploadPanelParentWidth() {        return viewRoot != null ? viewRoot.getWidth() : 0;    }    @Override    public int getUploadPanelParentHeight() {        return viewRoot != null ? viewRoot.getHeight() : 0;    }    @Override    public void getDataVideo(byte[] bytes, Camera.Size size) {        onCallVideoData(bytes, size);    }}