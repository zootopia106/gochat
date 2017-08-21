// Code generated by Wire protocol buffer compiler, do not edit.
// Source file: voip.proto
package red.tel.chat.generated_protobuf;

import android.os.Parcelable;
import com.squareup.wire.AndroidMessage;
import com.squareup.wire.FieldEncoding;
import com.squareup.wire.Message;
import com.squareup.wire.ProtoAdapter;
import com.squareup.wire.ProtoReader;
import com.squareup.wire.ProtoWriter;
import com.squareup.wire.WireField;
import com.squareup.wire.internal.Internal;
import java.io.IOException;
import java.lang.Integer;
import java.lang.Object;
import java.lang.Override;
import java.lang.String;
import java.lang.StringBuilder;
import okio.ByteString;

public final class AVQuality extends AndroidMessage<AVQuality, AVQuality.Builder> {
  public static final ProtoAdapter<AVQuality> ADAPTER = new ProtoAdapter_AVQuality();

  public static final Parcelable.Creator<AVQuality> CREATOR = AndroidMessage.newCreator(ADAPTER);

  private static final long serialVersionUID = 0L;

  public static final Integer DEFAULT_DIFF = 0;

  @WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer diff;

  public AVQuality(Integer diff) {
    this(diff, ByteString.EMPTY);
  }

  public AVQuality(Integer diff, ByteString unknownFields) {
    super(ADAPTER, unknownFields);
    this.diff = diff;
  }

  @Override
  public Builder newBuilder() {
    Builder builder = new Builder();
    builder.diff = diff;
    builder.addUnknownFields(unknownFields());
    return builder;
  }

  @Override
  public boolean equals(Object other) {
    if (other == this) return true;
    if (!(other instanceof AVQuality)) return false;
    AVQuality o = (AVQuality) other;
    return unknownFields().equals(o.unknownFields())
        && Internal.equals(diff, o.diff);
  }

  @Override
  public int hashCode() {
    int result = super.hashCode;
    if (result == 0) {
      result = unknownFields().hashCode();
      result = result * 37 + (diff != null ? diff.hashCode() : 0);
      super.hashCode = result;
    }
    return result;
  }

  @Override
  public String toString() {
    StringBuilder builder = new StringBuilder();
    if (diff != null) builder.append(", diff=").append(diff);
    return builder.replace(0, 2, "AVQuality{").append('}').toString();
  }

  public static final class Builder extends Message.Builder<AVQuality, Builder> {
    public Integer diff;

    public Builder() {
    }

    public Builder diff(Integer diff) {
      this.diff = diff;
      return this;
    }

    @Override
    public AVQuality build() {
      return new AVQuality(diff, super.buildUnknownFields());
    }
  }

  private static final class ProtoAdapter_AVQuality extends ProtoAdapter<AVQuality> {
    public ProtoAdapter_AVQuality() {
      super(FieldEncoding.LENGTH_DELIMITED, AVQuality.class);
    }

    @Override
    public int encodedSize(AVQuality value) {
      return ProtoAdapter.INT32.encodedSizeWithTag(1, value.diff)
          + value.unknownFields().size();
    }

    @Override
    public void encode(ProtoWriter writer, AVQuality value) throws IOException {
      ProtoAdapter.INT32.encodeWithTag(writer, 1, value.diff);
      writer.writeBytes(value.unknownFields());
    }

    @Override
    public AVQuality decode(ProtoReader reader) throws IOException {
      Builder builder = new Builder();
      long token = reader.beginMessage();
      for (int tag; (tag = reader.nextTag()) != -1;) {
        switch (tag) {
          case 1: builder.diff(ProtoAdapter.INT32.decode(reader)); break;
          default: {
            FieldEncoding fieldEncoding = reader.peekFieldEncoding();
            Object value = fieldEncoding.rawProtoAdapter().decode(reader);
            builder.addUnknownField(tag, fieldEncoding, value);
          }
        }
      }
      reader.endMessage(token);
      return builder.build();
    }

    @Override
    public AVQuality redact(AVQuality value) {
      Builder builder = value.newBuilder();
      builder.clearUnknownFields();
      return builder.build();
    }
  }
}
