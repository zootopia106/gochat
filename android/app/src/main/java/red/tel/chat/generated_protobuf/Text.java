// Code generated by Wire protocol buffer compiler, do not edit.
// Source file: wire.proto
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
import java.lang.Object;
import java.lang.Override;
import java.lang.String;
import java.lang.StringBuilder;
import okio.ByteString;

public final class Text extends AndroidMessage<Text, Text.Builder> {
  public static final ProtoAdapter<Text> ADAPTER = new ProtoAdapter_Text();

  public static final Parcelable.Creator<Text> CREATOR = AndroidMessage.newCreator(ADAPTER);

  private static final long serialVersionUID = 0L;

  public static final String DEFAULT_BODY = "";

  @WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#STRING"
  )
  public final String body;

  public Text(String body) {
    this(body, ByteString.EMPTY);
  }

  public Text(String body, ByteString unknownFields) {
    super(ADAPTER, unknownFields);
    this.body = body;
  }

  @Override
  public Builder newBuilder() {
    Builder builder = new Builder();
    builder.body = body;
    builder.addUnknownFields(unknownFields());
    return builder;
  }

  @Override
  public boolean equals(Object other) {
    if (other == this) return true;
    if (!(other instanceof Text)) return false;
    Text o = (Text) other;
    return unknownFields().equals(o.unknownFields())
        && Internal.equals(body, o.body);
  }

  @Override
  public int hashCode() {
    int result = super.hashCode;
    if (result == 0) {
      result = unknownFields().hashCode();
      result = result * 37 + (body != null ? body.hashCode() : 0);
      super.hashCode = result;
    }
    return result;
  }

  @Override
  public String toString() {
    StringBuilder builder = new StringBuilder();
    if (body != null) builder.append(", body=").append(body);
    return builder.replace(0, 2, "Text{").append('}').toString();
  }

  public static final class Builder extends Message.Builder<Text, Builder> {
    public String body;

    public Builder() {
    }

    public Builder body(String body) {
      this.body = body;
      return this;
    }

    @Override
    public Text build() {
      return new Text(body, super.buildUnknownFields());
    }
  }

  private static final class ProtoAdapter_Text extends ProtoAdapter<Text> {
    public ProtoAdapter_Text() {
      super(FieldEncoding.LENGTH_DELIMITED, Text.class);
    }

    @Override
    public int encodedSize(Text value) {
      return ProtoAdapter.STRING.encodedSizeWithTag(1, value.body)
          + value.unknownFields().size();
    }

    @Override
    public void encode(ProtoWriter writer, Text value) throws IOException {
      ProtoAdapter.STRING.encodeWithTag(writer, 1, value.body);
      writer.writeBytes(value.unknownFields());
    }

    @Override
    public Text decode(ProtoReader reader) throws IOException {
      Builder builder = new Builder();
      long token = reader.beginMessage();
      for (int tag; (tag = reader.nextTag()) != -1;) {
        switch (tag) {
          case 1: builder.body(ProtoAdapter.STRING.decode(reader)); break;
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
    public Text redact(Text value) {
      Builder builder = value.newBuilder();
      builder.clearUnknownFields();
      return builder.build();
    }
  }
}
